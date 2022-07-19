obj-m+=procesos.o

all:
	make -C /lib/modules/$(shell uname -r)/build/ M=$(PWD) modules
clean:
	make -C /lib/modules/$(shell uname -r)/build/ M=$(PWD) clean
run:
	sudo insmod procesos.ko
stop:
	sudo rmmod procesos.ko