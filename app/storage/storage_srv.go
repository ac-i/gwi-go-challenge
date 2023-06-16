package storage

import (
	"x-gwi/app/instance"
	"x-gwi/service"
)

type ServiceStorage struct {
	inst    *instance.Instance
	sstAQL  *ServiceStoreAQL
	sstKVBC *ServiceStoreKVBC
	name    service.CoreName
	noAQL   bool
}

func (st *ServiceStorage) InstMode() string {
	return st.inst.Mode()
}

func (st *ServiceStorage) InstName() string {
	return st.inst.Name()
}

func (st *ServiceStorage) CoreName() service.CoreName {
	return st.name
}

func (st *ServiceStorage) AQL() *ServiceStoreAQL {
	return st.sstAQL
}

func (st *ServiceStorage) KVBC() *ServiceStoreKVBC {
	return st.sstKVBC
}