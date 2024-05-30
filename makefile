.PHONY: heartbeat neuron

all: heartbeat neuron

heartbeat:
	go build -o output/heartbeat.exe ./heartbeat

neuron:
	go build -o output/neuron.exe ./neuron