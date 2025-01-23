import React from "react";
import { useEffect, useState } from "react";
import axios from "axios";
import { apiBaseUrl } from "../services/api";

interface Note {
  UserID: number;
  title: string;
  content: string;
  createdAt: string;
}

const Noteslist = () => {
  const [notes, setNotes] = useState<Note[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchNotes = async () => {
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          setError("Unauthorized. Please log in.");
          return;
        }

        const response = await axios.get(`${apiBaseUrl}/notes`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setNotes(response.data);
        console.log("Updated state:", response.data);
      } catch (err) {
        setError("Failed to fetch notes. Please try again.");
      }
    };

    fetchNotes();
  }, []);

  return (
    <>
      <div className="border-2">Noteslist</div>
      <div className="max-w-2xl mx-auto mt-8">
        <h1 className="text-2xl font-bold mb-4">Your Notes</h1>
        {error && <p className="text-red-500 mb-4">{error}</p>}
        {notes && notes.length ? (
  <div className="space-y-4">
    
    {notes.map((note, index) => (
  <div key={`${note.id}-${index}`} className="p-4 border rounded-lg shadow">
    <h2 className="text-xl font-semibold">{note.title}</h2>
    <p className="text-gray-600 mt-2">{note.content}</p>
    <p className="text-sm text-gray-400 mt-1">
      Created at: {new Date(note.createdAt).toLocaleString()}
    </p>
  </div>
))}


  </div>
) : (
  <p className="text-gray-500">No notes available.</p>
)}

      </div>
    </>
  );
};

export default Noteslist;
