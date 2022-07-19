import React from 'react';
import '../App.css';
import {useState} from "react";
import Tree from './proctree';
import api_host from './conf';


function Proc(){

    const [listproc, setlistproc] = React.useState({}
        /*[{pid:1, nombre:"hola", estado:"que tal", hijos:[{pid:3, nombre:"hola", estado:"que tal"}]},{pid:2, nombre:"hola", estado:"que tal"}]*/
    );

const getProc = async(evt) => {
    evt.preventDefault();
    var name = evt.target.name;
    var endpoint = "";
    if (name === "p1"){
        endpoint = `${api_host.api_one}` + evt.target.value;
    } else { 
        endpoint = `${api_host.api_two}` + evt.target.value;
    }
    var value = null;
    while(value == null){
        const response =
          await fetch(endpoint,
          { headers: {'Content-Type': 'application/json'} }
        )
        var maquina = await response.json();
        if (name === "p1" && maquina.mv == 'vm1'){
            value = maquina;
        } else if (name === "p2" && maquina.mv == 'vm2'){ 
            value = maquina;
        }
        
    }
        setlistproc(value);
    
}
 

    
    return(
        <div class="box">
            <button class="button is-link" onClick={getProc} name="p1" value="procesos1">PROCESOS VM1</button>
            <button class="button is-danger" onClick={getProc} name="p2" value="procesos2">PROCESOS VM2</button>
            <p class="title is-1">{listproc.mv}</p>
            <div class="table-container">
                <table class="table is-bordered is-striped is-hoverable is-fullwidth is-warning">
                    <thead>
                
                        <th>  <div class="notification is-warning">PID </div></th>
                        <th><div class="notification is-warning">NOMBRE</div></th>
                        <th ><div class="notification is-warning">ESTADO</div></th>
                        
                    </thead>
                        <Tree data={listproc.data} />
                    </table>
                
                
        </div>
        </div>
        
    );
    

}



export default Proc;