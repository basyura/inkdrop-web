import { useEffect } from "react";
import { booksAtom, currentBookAtom } from "../atoms/Atoms";
import { useAtom } from "jotai";
import webapi from "../api/webapi";

const Books = () => {
  const [books, setBooks] = useAtom(booksAtom);
  const [, setCurrentBook] = useAtom(currentBookAtom);

  useEffect(() => {
    async function fetchData() {
      const res = await webapi.get("books");
      console.log("Books#useEffect", res.data);
      setBooks(res.data);
      setCurrentBook(res.data[0]);
    }
    fetchData();
  }, []);
  return (
    <div>
      <ul>
        {books.map((v) => (
          <li key={v._id}>{v.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default Books;
