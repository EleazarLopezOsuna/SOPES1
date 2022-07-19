
import './App.css';
import Sock from './components/sock';
import React, {useState} from "react";

function App() {
  const [loadClient, setLoadClient] = useState(false);
  return (
    <div class="box">
      {/* LOAD OR UNLOAD THE CLIENT */}
      <button class="button is-success" onClick={() => setLoadClient(prevState => !prevState)}>
        START CLIENT
      </button>
      {/* SOCKET IO CLIENT*/}
      {loadClient ? <Sock /> : null}
    </div>
  );
}


export default App;
