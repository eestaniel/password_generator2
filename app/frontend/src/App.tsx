import {useState} from 'react';
import './App.css';
import {Greet, GeneratePassword} from "../wailsjs/go/main/App";

function App() {
    const [password, setPassword] = useState("")
    const [length, setLength] = useState(12)
    const lengthSlider = document.getElementById("length_slider") as HTMLInputElement
    const lengthInput = document.getElementById("length_input")as HTMLInputElement;


    function generatePassword(e: { preventDefault: () => void; }) {
        GeneratePassword(length).then((res: any) => {
            setPassword(res)
        })
        e.preventDefault()
    }

    function handleLengthChange(e: { target: { value: string; }; }) {
        setLength(parseInt(e.target.value))
        lengthSlider.value = e.target.value
        lengthInput.value = e.target.value
    }

    return (
        <div id="App">
            <div className="pw-header">
                <h1>Password Generator</h1>
            </div>
            <form className="pw-form">
                <div className="pw-form-item">
                    <input type={"text"} placeholder={"Password"} value={password} readOnly={true}/>
                </div>
                <div className="pw-form-other">
                    <div className="checkbox">
                        <input type="checkbox" id={"uppercase"}/>
                        <label htmlFor={"uppercase"}>uppercase</label>
                    </div>
                    <div className="checkbox">
                        <input type="checkbox" id={"lowercase"}/>
                        <label htmlFor={"lowercase"}>lowercase</label>
                    </div>
                    <div className="checkbox">
                        <input type="checkbox" id={"number"}/>
                        <label htmlFor={"number"}>number</label>
                    </div>
                    <div className="checkbox">
                        <input type="checkbox" id={"special"}/>
                        <label htmlFor={"special"}>special</label>
                    </div>
                    <div className="slider">
                        <input type="range" id={"length_slider"} min={"1"} max={"256"} value={length} onChange={handleLengthChange}/>
                        <label htmlFor={"length"}></label>
                        <input type="number" id={"length_input"} min={"1"} max={"256"} value={length} onChange={handleLengthChange}/>
                        <label htmlFor={"length_input"}></label>
                    </div>
                </div>
                <button onClick={generatePassword}>Generate Password</button>
            </form>
        </div>
    )
}

export default App
