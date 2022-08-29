import { currentNoteAtom } from "../atoms/Atoms";
import { useAtomValue } from "jotai";
import MarkdownIt from "markdown-it";

const Note = () => {
  const currentNote = useAtomValue(currentNoteAtom);
  const md = new MarkdownIt();
  if (currentNote.title == null) {
    return <div></div>;
  }
  const RenderHTML = (props) => (
    <div dangerouslySetInnerHTML={{ __html: props.HTML }}></div>
  );
  return (
    <div>
      <h1> {currentNote.title}</h1>
      <RenderHTML HTML={md.render(currentNote.body)} />
    </div>
  );
};

export default Note;
