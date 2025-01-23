import React, { useState } from 'react'
import notesImg from "../../public/notes.png"
import axios from 'axios';
import { apiBaseUrl } from '../services/api';

const AddNote = () => {
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");
    const [error, setError] = useState("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("");
    
        try {
          const token = localStorage.getItem("token");
          if (!token) {
            setError("Unauthorized. Please log in.");
            return;
          }
    
          await axios.post(
            `${apiBaseUrl}/notes`,
            { title, content },
            {
              headers: {
                Authorization: `Bearer ${token}`,
              },
            }
          );
    
          setTitle("");
          setContent("");
          alert("Note added successfully!");
        } catch (err) {
          setError("Failed to add note. Please try again.");
        }
      };

      
  return (
    <div className='border-2 min-h-screen m-0'>
        <img src={notesImg} alt="" className='h-[25rem] border-2 mx-auto' />
        <div className="max-w-md mx-auto mt-8">
      <h1 className="text-2xl font-bold mb-4">Add Note</h1>
      {error && <p className="text-red-500 mb-4">{error}</p>}
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="input input-bordered w-full mb-4"
          required
        />
        <textarea
          placeholder="Content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="textarea textarea-bordered w-full mb-4"
          required
        />
        <button type="submit" className="btn btn-primary w-full">
          Add Note
        </button>
      </form>
    </div>
    </div>
  )
}

export default AddNote