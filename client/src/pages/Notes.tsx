import React from 'react';
import Sidebar from '../components/Sidebar';
import { useState, useEffect } from 'react';
import AddNote from '../components/AddNote';
import Noteslist from '../components/Noteslist';


const Notes = () => {

  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <Sidebar />

      {/* Main content area with left margin equal to sidebar width */}
      <div className="ml-64 flex-1">
        <nav className="h-20 bg-black font-sans"></nav>
        <div className="p-4 flex">
          <div className='w-[50%] h-screen'><Noteslist/></div>
          <div className='w-[50%]'><AddNote/></div>
        </div>
      </div>
    </div>
  );
};

export default Notes;
