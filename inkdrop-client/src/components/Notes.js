import { useEffect } from "react";
import { notesAtom, currentNoteAtom } from "../atoms/Atoms";
import { useAtom } from "jotai";
import webapi from "../api/webapi";

const Notes = () => {
  const [notes, setNotes] = useAtom(notesAtom);
  const [, setCurrentNote] = useAtom(currentNoteAtom);

  useEffect(() => {
    async function fetchData() {
      const res = await webapi.get("notes");
      console.log("Notes#useEffect", res.data);
      setNotes(res.data);
    }
    fetchData();
  }, []);

  const handleSelectNote = (note) => {
    console.log(note.title);
    setCurrentNote(note);
  };

  return (
    <div className={"notes"}>
      <ul>
        {notes.map((v) => (
          <li key={v._id} onClick={() => handleSelectNote(v)}>
            {v.title}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Notes;
