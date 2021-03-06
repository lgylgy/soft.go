package ecore

type eNotifyingListInternal interface {
	ENotifyingList

	inverseAdd(object interface{}, chain ENotificationChain) ENotificationChain

	inverseRemove(object interface{}, chain ENotificationChain) ENotificationChain
}

// ENotifyingListImpl ...
type ENotifyingListImpl struct {
	*basicEList
}

// NewENotifyingListImpl ...
func NewENotifyingListImpl() *ENotifyingListImpl {
	l := new(ENotifyingListImpl)
	l.basicEList = NewUniqueBasicEList([]interface{}{})
	l.interfaces = l
	return l
}

func newENotifyingListImplFromData(data []interface{}) *ENotifyingListImpl {
	l := new(ENotifyingListImpl)
	l.basicEList = NewUniqueBasicEList(data)
	l.interfaces = l
	return l
}

// GetNotifier ...
func (list *ENotifyingListImpl) GetNotifier() ENotifier {
	return nil
}

// GetFeature ...
func (list *ENotifyingListImpl) GetFeature() EStructuralFeature {
	return nil
}

// GetFeatureID ...
func (list *ENotifyingListImpl) GetFeatureID() int {
	return -1
}

type notifyingListNotification struct {
	*abstractNotification
	list *ENotifyingListImpl
}

func (notif *notifyingListNotification) GetNotifier() ENotifier {
	return notif.list.interfaces.(ENotifyingList).GetNotifier()
}

func (notif *notifyingListNotification) GetFeature() EStructuralFeature {
	return notif.list.interfaces.(ENotifyingList).GetFeature()
}

func (notif *notifyingListNotification) GetFeatureID() int {
	return notif.list.interfaces.(ENotifyingList).GetFeatureID()
}

func (list *ENotifyingListImpl) createNotification(eventType EventType, oldValue interface{}, newValue interface{}, position int) ENotification {
	n := new(notifyingListNotification)
	n.abstractNotification = NewAbstractNotification(eventType, oldValue, newValue, position)
	n.interfaces = n
	n.list = list
	return n
}

func (list *ENotifyingListImpl) isNotificationRequired() bool {
	notifier := list.interfaces.(ENotifyingList).GetNotifier()
	return notifier != nil && notifier.EDeliver() && !notifier.EAdapters().Empty()
}

func (list *ENotifyingListImpl) createAndAddNotification(ns ENotificationChain, eventType EventType, oldValue interface{}, newValue interface{}, position int) ENotificationChain {
	notifications := ns
	if list.isNotificationRequired() {
		notification := list.createNotification(eventType, oldValue, newValue, position)
		if notifications != nil {
			notifications.Add(notification)
		} else {
			notifications = notification.(ENotificationChain)
		}
	}
	return notifications
}

func (list *ENotifyingListImpl) createAndDispatchNotification(notifications ENotificationChain, eventType EventType, oldValue interface{}, newValue interface{}, position int) {
	list.createAndDispatchNotificationFn(notifications, func() ENotification {
		return list.createNotification(eventType, oldValue, newValue, position)
	})
}

func (list *ENotifyingListImpl) createAndDispatchNotificationFn(notifications ENotificationChain, createNotification func() ENotification) {
	if list.isNotificationRequired() {
		notification := createNotification()
		if notifications != nil {
			notifications.Add(notification)
			notifications.Dispatch()
		} else {
			notifier := list.interfaces.(ENotifyingList).GetNotifier()
			if notifier != nil {
				notifier.ENotify(notification)
			}
		}
	} else {
		if notifications != nil {
			notifications.Dispatch()
		}
	}
}

func (list *ENotifyingListImpl) inverseAdd(object interface{}, notifications ENotificationChain) ENotificationChain {
	return notifications
}

func (list *ENotifyingListImpl) inverseRemove(object interface{}, notifications ENotificationChain) ENotificationChain {
	return notifications
}

// AddWithNotification ...
func (list *ENotifyingListImpl) AddWithNotification(object interface{}, notifications ENotificationChain) ENotificationChain {
	index := list.Size()
	list.basicEList.doAdd(object)
	return list.createAndAddNotification(notifications, ADD, nil, object, index)
}

// RemoveWithNotification ...
func (list *ENotifyingListImpl) RemoveWithNotification(object interface{}, notifications ENotificationChain) ENotificationChain {
	index := list.IndexOf(object)
	if index != -1 {
		oldObject := list.basicEList.RemoveAt(index)
		return list.createAndAddNotification(notifications, REMOVE, oldObject, nil, index)
	}
	return notifications
}

// SetWithNotification ...
func (list *ENotifyingListImpl) SetWithNotification(index int, object interface{}, notifications ENotificationChain) ENotificationChain {
	oldObject := list.basicEList.doSet(index, object)
	return list.createAndAddNotification(notifications, SET, oldObject, object, index)
}

func (list *ENotifyingListImpl) doAdd(object interface{}) {
	index := list.Size()
	list.basicEList.doAdd(object)
	notifications := list.interfaces.(eNotifyingListInternal).inverseAdd(object, nil)
	list.createAndDispatchNotification(notifications, ADD, nil, object, index)
}

func (list *ENotifyingListImpl) doAddAll(l EList) bool {
	return list.doInsertAll(list.Size(), l)
}

func (list *ENotifyingListImpl) doInsert(index int, object interface{}) {
	list.basicEList.doInsert(index, object)
	notifications := list.interfaces.(eNotifyingListInternal).inverseAdd(object, nil)
	list.createAndDispatchNotification(notifications, ADD, nil, object, index)
}

func (list *ENotifyingListImpl) doInsertAll(index int, l EList) bool {
	if l.Empty() {
		return false
	}

	result := list.basicEList.doInsertAll(index, l)
	var notifications ENotificationChain = NewNotificationChain()
	for it := l.Iterator(); it.HasNext(); {
		notifications = list.interfaces.(eNotifyingListInternal).inverseAdd(it.Next(), notifications)
	}
	list.createAndDispatchNotificationFn(notifications, func() ENotification {
		if l.Size() == 1 {
			return list.createNotification(ADD, nil, l.Get(0), index)
		} else {
			return list.createNotification(ADD_MANY, nil, l.ToArray(), index)
		}
	})
	return result
}

func (list *ENotifyingListImpl) doSet(index int, newObject interface{}) interface{} {
	oldObject := list.basicEList.doSet(index, newObject)
	if newObject != oldObject {
		var notifications ENotificationChain
		notifications = list.interfaces.(eNotifyingListInternal).inverseRemove(oldObject, notifications)
		notifications = list.interfaces.(eNotifyingListInternal).inverseAdd(newObject, notifications)
		list.createAndDispatchNotification(notifications, SET, oldObject, newObject, index)
	}
	return oldObject
}

func (list *ENotifyingListImpl) doClear() []interface{} {
	oldData := list.basicEList.doClear()
	if len(oldData) == 0 {
		list.createAndDispatchNotification(nil, REMOVE_MANY, list, nil, -1)
	} else {
		var notifications ENotificationChain = NewNotificationChain()
		for _, e := range oldData {
			notifications = list.interfaces.(eNotifyingListInternal).inverseRemove(e, notifications)
		}

		list.createAndDispatchNotificationFn(notifications,
			func() ENotification {
				if list.Size() == 1 {
					return list.createNotification(REMOVE, nil, list.Get(0), 0)
				} else {
					return list.createNotification(REMOVE_MANY, nil, list.ToArray(), -1)
				}
			})
	}
	return oldData
}

// RemoveAt ...
func (list *ENotifyingListImpl) RemoveAt(index int) interface{} {
	oldObject := list.basicEList.RemoveAt(index)
	var notifications ENotificationChain
	notifications = list.interfaces.(eNotifyingListInternal).inverseRemove(oldObject, notifications)
	list.createAndDispatchNotification(notifications, REMOVE, oldObject, nil, index)
	return oldObject
}
