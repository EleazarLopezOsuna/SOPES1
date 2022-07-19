import React, {useState, useEffect } from "react";
import socketIOClient from "socket.io-client";
import '../App.css';
import Tree from './proctree';
import api_host from './conf';

function Sock(){
    const [Ram, setRam] = useState([]);
    const [Procs, setProcs] = useState([]);

    useEffect(()=> {
        const socket = socketIOClient(`${api_host.api_sock}`);
        socket.on("ramLog", data => {
            setRam(data);
        });
        socket.on("procLog", data => {
            setProcs(data);
        });
        return () => socket.disconnect();
    }, []);

    return (
        <>
        <div class="box">
        <p class="title is-1">LOG RAM</p>
        <div class="table-container">
            <table class="table is-bordered is-striped is-hoverable is-fullwidth is-warning">
                <thead>    
                    <th>  <div class="notification is-warning">ID </div></th>
                    <th><div class="notification is-warning">NOMBRE VM</div></th>
                    <th><div class="notification is-warning">ENDPOINT</div></th>
                    <th ><div class="notification is-warning">DATE</div></th>
                      
                </thead>
                <LogDataRam data={Ram} /> 
            </table>
        </div>
        </div>



        <div class="box">
        <p class="title is-1">LOG PROCESOS</p>
        <div class="table-container">
            <table class="table is-bordered is-striped is-hoverable is-fullwidth is-warning">
                <thead>    
                    <th>  <div class="notification is-warning">ID </div></th>
                    <th><div class="notification is-warning">NOMBRE VM</div></th>
                    <th><div class="notification is-warning">ENDPOINT</div></th>
                    <th ><div class="notification is-warning">DATE</div></th>
                    <th ><div class="notification is-warning">PROCESOS</div></th>
                      
                </thead>
                <LogDataProc data={Procs} />
                
            </table>
        </div>
        </div>
    </>
    );
}

const LogDataRam = ({ data = [] }) => {

    const [hijoVisible, sethijoVisible] = useState(false);
    const hasChild = true;
    return (
        <>
       
          {data.map((tree) => (
              <>
                <tbody>
                    <tr onClick={e => sethijoVisible((v) => !v)}>
                        <th> {tree._id}</th>   
                        <td>{tree.nombrevm}</td>     
                        <td>{tree.endpoint}</td>
                        <td>{tree.date}</td>
                    </tr>
                </tbody>

            
                {hasChild && hijoVisible && (
                    <tr >
                        <th COLSPAN="4">
                    
                            <div class="box">
                                <div class="columns">
                                
                                    <div class="column">
                                    <div class="notification is-info">
                                        MemoriaTotal: {tree.data[0].total}        
                                        </div>       
                                    </div>
                                    <div class="column">
                                    <div class="notification is-info">
                                        MemoriaUtilizada: {tree.data[0].memoriaenuso}
                                        </div>
                                        </div>
                                    <div class="column">
                                    <div class="notification is-info">
                                        Porcentaje: {tree.data[0].porcentaje}
                                        </div>
                                        </div>
                                    <div class="column">
                                    <div class="notification is-info">
                                         MemoriaLibre: {tree.data[0].memorialibre}
                                         </div>
                                         </div>
                                </div>
                            </div>
                        </th>
                    </tr>
                )}
                </>
            ))}
      </>
    );

    
  };

  const ProcLogRow =({data = [ ]}) => {
    const [hijoVisible, sethijoVisible] = useState(false);
    const hasChild = true;
    return (
        <>
       
         
              <>
              <tr >
                <th> {data._id}</th>   
                <td>{data.nombrevm}</td>     
                <td>{data.endpoint}</td>
                <td>{data.date}</td>
                <td><button class="button is-info is-fullwidth" onClick={e => sethijoVisible((v) => !v)}>PROCESOS</button></td>
                
                </tr>
                {hasChild && hijoVisible && (
                    <tr><th COLSPAN="5"><ProcsLog data={data.procs}/></th></tr>
                )}
            </>
          
        
      </>
    );
  }


  const LogDataProc = ({ data = [] }) => {
    return (
        <>
       
          {data.map((tree) => (
                <ProcLogRow data={tree} />
          ))}
        
      </>
    );

}


const ProcsLog = ({data =[ ]}) => {
    return (
        
        <>
            
                
                    
                <section class="section">
                <div class="table-container">
                    <table class="table is-bordered is-striped is-hoverable is-fullwidth is-warning">
                        <thead>
                
                            <th>  <div class="notification is-warning">PID </div></th>
                            <th><div class="notification is-warning">NOMBRE</div></th>
                            <th ><div class="notification is-warning">ESTADO</div></th>
                        
                        </thead>
                        
                            <Tree data={data} />
                        
                    </table>
                
                
                </div>
        
                </section>
            
        </>

    );
}

export default Sock;