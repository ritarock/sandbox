import {useState} from 'react';
import './App.css';
import {Generate} from "../wailsjs/go/main/App";

function App() {
    const [digits, setDigits] = useState("")
    const [pass, setPass] = useState("")
    const [err, setErr] = useState("")
    const updatedigits = (e: any) => setDigits(e.target.value);

    function generate() {
        Generate(digits)
        .then((result) => {
            setPass(result);
        })
        .catch((error) => {
            setErr(error);
        })
    }

    return (
        <div id="App">
            <div>{err}</div>
            <div id="input" className="input-box">
                <input id="digits" className="input" onChange={updatedigits} autoComplete="off" type="text" />
                <button className="btn" onClick={generate}>generate</button>
                <p>{pass}</p>
            </div>
        </div>
    )
}

export default App
