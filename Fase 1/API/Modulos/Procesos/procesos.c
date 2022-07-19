#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/sched/signal.h>
#include <linux/sched.h>
#include <linux/fs.h>
#include <linux/version.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Sopes1_G2");

int contadorPadre;
int contadorPadreAux;
int contadorHijo;
int contadorHijoAux;
struct task_struct *task; // estructura definida en sched.h para tareas/procesos
struct task_struct *childtask; // estructura necesaria para iterar a travez de procesos secundarios
struct list_head *list; // estructura necesaria para recorrer cada lista de tareas tarea->estructura de hijos

void escribir_hijos(struct seq_file *archivo, struct task_struct *task){
    contadorHijo = 0;
    contadorHijoAux = 1;

    list_for_each( list,&task->children ){
        contadorHijo++;
    }

    seq_printf(archivo, "\"hijo\":[");

        list_for_each( list,&task->children ){
            seq_printf(archivo, "\t{\n");
            childtask = list_entry( list, struct task_struct, sibling );
            seq_printf(archivo, "\t\"pid\": \"%d\",\n",childtask->pid);
            seq_printf(archivo, "\t\"nombre\": \"%s\",\n",childtask->comm);
            seq_printf(archivo, "\t\"estado\": \"%ld\" \n",childtask->state);
            if (contadorHijoAux == contadorHijo){
                seq_printf(archivo, "\t}");
            }else{
                seq_printf(archivo, "\t},\n");
            }
            
            contadorHijoAux++;
        }

    seq_printf(archivo, "]\n");
}

static int escribir_archivo(struct seq_file * archivo, void *v){
    contadorPadre = 0;
    contadorPadreAux = 1;

    for_each_process( task ) {
       contadorPadre++;
    }

    seq_printf(archivo, "[");
    for_each_process( task ){
        seq_printf(archivo, "{\n");
        seq_printf(archivo, "\"pid\": \"%d\",\n",task->pid);
        seq_printf(archivo, "\"nombre\": \"%s\",\n",task->comm);
        seq_printf(archivo, "\"estado\": \"%ld\",\n",task->state);

        escribir_hijos(archivo,task);

        if (contadorPadreAux == contadorPadre){
            seq_printf(archivo, "}]\n");
        }else{
            seq_printf(archivo, "},\n");
        }
        
        contadorPadreAux++;
    }

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

int iniciar(void){ //modulo de inicio
	proc_create("cpu_grupo2", 0, NULL, &operaciones);
	printk(KERN_INFO "%s", "“Módulo lista de procesos del Grupo 2 Cargado");

	return 0;
}

void salir(void){
	remove_proc_entry("cpu_grupo2",NULL);
	printk(KERN_INFO "%s","Módulo lista de procesos del Grupo 2 Desmontado");	
}

module_init(iniciar);
module_exit(salir);