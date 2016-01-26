package main

type ServiceList map[string]Service

type Service struct {
	ID      string `json:"ID"`
	Service string `json:"Service"`
	Address string `json:"Address"`
	Port    int    `json:"Port"`
}
