#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>
#include <linux/hugetlb.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/fs.h>

#define  BUFSIZE 150

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Sopes1_G2");

long memoriaTotal;
long memoriaLibre;
struct sysinfo inf;

static int escribir_archivo(struct seq_file * archivo, void *v){
    si_meminfo(&inf);
    memoriaTotal = (inf.totalram * 4);
    memoriaLibre = (inf.freeram * 4);
    seq_printf(archivo, "{\n");
    seq_printf(archivo, "\"memoriaTotal\":\"%8lu\",\n",memoriaTotal/1024);
    seq_printf(archivo, "\"memoriaLibre\":\"%8lu\",\n",memoriaLibre/1024);
    seq_printf(archivo, "\"memoriaUso\":\"%8lu\"\n",(memoriaTotal-memoriaLibre)/1024);
    seq_printf(archivo, "}\n");
	return 0;
}

static int al_abrir(struct inode *inode, struct file*file){
	return single_open(file, escribir_archivo, NULL);	
}

static ssize_t write_proc(struct file *file, const char __user *bufer, size_t count, loff_t *offp){
    return 0;
}

static struct proc_ops operaciones = {
    .proc_open = al_abrir,
    .proc_read = seq_read,
    .proc_write = write_proc,
	.proc_release = single_release,
};

static int iniciar(void){ //modulo de inicio
	proc_create("ram_grupo2", 0, NULL, &operaciones);
	printk(KERN_INFO "%s", "Módulo RAM del Grupo2 Cargado");

 	return 0;
}

static void salir(void){
	remove_proc_entry("ram_grupo2",NULL);
	printk(KERN_INFO "%s","Módulo RAM del Grupo2 Desmontado");
	
}

module_init(iniciar);
module_exit(salir);