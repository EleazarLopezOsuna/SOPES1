obj-m+=ram.o

all:
	make -C /lib/modules/$(shell uname -r)/build/ M=$(PWD) modules
clean:
	make -C /lib/modules/$(shell uname -r)/build/ M=$(PWD) clean
run:
	sudo insmod ram.ko
stop:
	sudo rmmod ram.ko