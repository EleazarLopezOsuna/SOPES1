import React, {useState, useEffect} from "react";
import socketIOClient from "socket.io-client";
import Ultimo from './results/last';
import Mejor from './results/best';
import '../App.css';
import api_host from './conf';
function Tidis(){
    const [vLast, setVLast] = useState(false);
    const [vBest, setVBest] = useState(false);

    
    const [logList, setLogList] = useState(['{game_id:2,players:100,game_name:Game1,winner:8,queue:Kafka}']);

    useEffect(()=> {
        const socket = socketIOClient(`${api_host.api_sock}`);
        socket.on("tidis", data => {
            setLogList(data);
        });
        return () => socket.disconnect();
    }, []);

    

    return(
        <>

            <div class="box">
                <h1 class="title">TIDIS</h1>
                <div class="buttons">
                
                <button class="button is-success" onClick={() => setVLast(prevState => !prevState)}>
                    Ultimos 10 Juegos
                </button>
                <button class="button is-success" onClick={() => setVBest(prevState => !prevState)}>
                    Jugadores
                </button>
                </div>
            
                {/*window.alert     (     JSON.stringify(typeof( logList.res[0]) )     )*/}
                {vLast ? <div class="box"><Ultimo data={logList}/></div> : null}
                
                {vBest ?  <div class="box"><Mejor data={logList}/></div> : null}
                
            </div>
        </>
    );
}

export default Tidis;