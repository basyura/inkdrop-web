import { currentNoteAtom } from "../atoms/Atoms";
import { useAtomValue } from "jotai";

const Note = () => {
  const currentNote = useAtomValue(currentNoteAtom);
  return (
    <div>
      <h1> {currentNote.title}</h1>
      <div> {currentNote.body} </div>
    </div>
  );
};

export default Note;
