import React, {useState} from "react";
import axios from "axios"
import '../App.css';
import api_host from './conf';
import Games from './graflog/ggame';
import Broker from './graflog/gbroker';
import Tabla from './graflog/tabla';


function Mongo(){
    const [loadBroker, setLoadBroker] = useState(false);
    const [loadJuego, setLoadJuego] = useState(false);
    const [loadTable, setLoadTable] = useState(false);
    const [logList, setLogList] = useState([]);

    const getLogs = async (evt) => {
        evt.preventDefault();
        const endpoint = `http://35.223.10.234:8080/get-all`;

        axios({
            method: 'get',
            url: endpoint,
          })
            .then(function (response) {
              setLogList(response.data);
            });

        /*
        const response =
            await fetch(endpoint, { headers: {'Content-Type': 'application/json'} } )
        var value = await response.json();
        console.log(value)
        */
        //setLogList(value);
    }


    return(
        <>
            <div class="box">
            <h1 class="title">LOGS MONGO</h1>
                <div class="buttons">
                <button class="button is-info is-large" onClick={getLogs}>
                    Get Logs
                </button>
                <button class="button is-success" onClick={() => setLoadBroker(prevState => !prevState)}>
                    Brokers
                </button>
                <button class="button is-success" onClick={() => setLoadJuego(prevState => !prevState)}>
                    Juegos
                </button>
                <button class="button is-success" onClick={() => setLoadTable(prevState => !prevState)}>
                    Logs
                </button>
                </div>
                {loadBroker ? <div class="box"><Broker logs={logList}/></div> : null}
                {loadJuego ?  <div class="box"><Games logs={logList}/></div> : null}
                {loadTable ?  <div class="box"><Tabla data={logList}/></div> : null}
            </div>
        </>
    );

}


export default Mongo;