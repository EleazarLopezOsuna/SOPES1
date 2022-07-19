import {React, useState, useEffect} from 'react';
import '../../App.css'

function Ultimo ({ data = []}) {
    
    
    
    return (
        <><h1 class="title">Ultimos Juegos</h1>
            <div class="table-container">
                <table class="table is-bordered is-striped is-hoverable is-fullwidth is-warning">
                    <thead>
                
                        <th>  <div class="notification is-warning">GAME ID </div></th>
                        <th><div class="notification is-warning">PLAYERS</div></th>
                        <th ><div class="notification is-warning">GAME NAME </div></th>
                        <th><div class="notification is-warning">WINNER</div></th>
                        <th ><div class="notification is-warning">QUEUE</div></th>
                        
                    </thead>
                        <List data={data} />
                    </table>
          
            </div>
      </>
    );
}

const List = ({data = []}) => {
    let lista = []
    try {
    data.map((node) => {
        lista.push(JSON.parse(node))
    })
    lista.reverse()
    lista = lista.slice(0,10)
    }catch(err)  {
        //window.alert(err)
    }

     return (
        <>
            {lista.map((node) => (

                <tbody>
                <tr>
                            <td> {node.game_id}</td>   
                            <td>{node.players}</td>     
                            <td>{node.game_name}</td>
                            <td>{node.winner}</td>     
                            <td>{node.queue}</td>
                </tr>
                </tbody>
            ))}
        </>
    );
}

export default Ultimo;

