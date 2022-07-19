import React, {useState} from "react";
import  './App.css';
import Mongo from './components/mongo.js';
import Redis from './components/redis.js';
import Tidis from './components/tidis.js';


function App() {
    const [logs, setLogs] = useState(true);
    const [vRedis, setVRedis] = useState(false);
    const [vTidis, setVTidis] = useState(false);

    const seeLogs = async (evt) => {
      setLogs(true);
      setVRedis(false);
      setVTidis(false);
    }

    const seeRedis = async (evt) => {
      setLogs(false);
      setVRedis(true);
      setVTidis(false);
    }

    const seeTidis = async (evt) => {
      setLogs(false);
      setVRedis(false);
      setVTidis(true);
    }

  return (
    <>
      <nav class="navbar is-dark" role="navigation" aria-label="main navigation">
        <div  class="navbar-menu">
          <div class="navbar-start">
            <a class="navbar-item" onClick={seeLogs}>
              Logs
            </a>
            <a class="navbar-item" onClick={seeRedis}>
              Redis
            </a>
            <a class="navbar-item" onClick={seeTidis}>
              Tidis
            </a>
          </div>
        </div>
      </nav>
      {logs ? <Mongo />: null} 
      {vRedis ? <Redis />: null} 
      {vTidis ? <Tidis />: null} 
    </>
  );
}

export default App;
