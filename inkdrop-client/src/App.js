import "./App.css";
import Books from "./components/Books";
import Notes from "./components/Notes";
import Note from "./components/Note";

function App() {
  return (
    <div className="App">
      <div className="sidebar">
        <Books />
      </div>
      <div className="notesContainer">
        <Notes />
      </div>
      <div className="noteContainer">
        <Note />
      </div>
    </div>
  );
}

export default App;
