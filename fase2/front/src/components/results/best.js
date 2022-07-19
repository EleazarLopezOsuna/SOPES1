import {React, useState} from 'react';
import '../../App.css'
import Tabla from '../graflog/tabla'

function Mejor ({ data = []}){

    let lista = []
    try {
    data.map((node) => {
        lista.push(JSON.parse(node))
    })
    var ganadores = {}

    lista.map(
        (juego) => {
            if(ganadores[juego.winner] == (null || undefined) ){
                ganadores[juego.winner] = {jugador: juego.winner ,vic :[juego]}
            } else {
                ganadores[juego.winner].vic.push(juego)
            }
        }
    )

    }catch  {
        
}
    return (
        <>
        <h1 class="title">Mejores Jugadores</h1>
            <div class="table-container">
                <table class="table is-bordered is-striped is-hoverable is-fullwidth is-warning">
                    <thead>
                
                        <th>  <div class="notification is-warning">JUGADOR </div></th>
                        <th><div class="notification is-warning">VICTORIAS</div></th>
                        
                        
                    </thead>
                        <List data={ganadores} />
                    </table>
          
            </div>
      </>
    );
}

function List ({ data = {} }) {
    let toList = []
        for (let val in data){
            toList.push(data[val])
            }
    toList.sort ( (a, b) => {
        if (a.vic.length > b.vic.length) {
            return -1;
          }
          if (a.vic.length < b.vic.length) {
            return 1;
          }
          // a debe ser igual b
          return 0;
    })
    toList = toList.slice(0,10)
    return (
        <>
            {toList.map((nodo) => (
               <Jugador data={nodo} />
            
               ))}
        
      </>
    );
  };

const Jugador = ({data = {}}) => {
    const [hijoVisible, sethijoVisible] = useState(false);

     return (
        <>
            
                <tbody>
                <tr onClick={e => sethijoVisible((v) => !v)}>
                    
                            <th> {data.jugador}</th>   
                            <td>{data.vic.length}</td>     
                </tr>
                </tbody>
                { hijoVisible && (
                <tr><th COLSPAN="2"><Tabla data={data.vic} /></th></tr>
        )}
        </>
    );
}

export default Mejor;

