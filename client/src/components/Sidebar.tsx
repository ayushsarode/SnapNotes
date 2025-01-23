import { CiStickyNote } from "react-icons/ci";
import { FaTasks } from "react-icons/fa";
import { IoLogOutOutline } from "react-icons/io5";
import { Link } from "react-router-dom";

const Sidebar = () => {
  return (
    <div className="">
      <aside
        id="default-sidebar"
        className="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0"
        aria-label="Sidebar"
      >
        <div className="h-full mt-[5rem] px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
          <p className="text-gray-500 px-2 text-sm tracking-wider py-3">
            WORKSPACE
          </p>
          <ul className="space-y-2 font-medium">
            <li>
              <Link
                to="/dashboard/notes"
                className="flex hover:bg-orange-300 hover:text-white transition duration-250 items-center py-3 p-2  text-black font-medium rounded-lg text-md group"
              >
                <CiStickyNote className="flex-shrink-0 w-5 scale-125 transition text-black stroke-[1] duration-75 group-hover:text-white" />
                <span className="ms-3">Notes</span>
              </Link>
            </li>

            <li>
              <Link
                to="/dashboard/tasks"
                className="flex hover:bg-orange-300 hover:text-white transition duration-250 items-center py-3 p-2 text-black font-medium rounded-lg text-md group"
              >
                <FaTasks className="flex-shrink-0 w-5 scale-125 transition text-black stroke-[1] duration-75 group-hover:text-white" />
                <span className="ms-3">Task</span>
              </Link>
            </li>

            <li>
              <Link
                to="#"
                className="flex hover:text-white transition duration-250 items-center p-2 text-white text-lg bg-red-500 hover:bg-red-600 rounded-lg text-md group text-start ml-0"
              >
                <IoLogOutOutline className="flex-shrink-0 w-5 scale-125 transition stroke-[1] duration-75 group-hover:text-white" />
                <span className="ms-3">Log Out</span>
              </Link>
            </li>
          </ul>
        </div>
      </aside>
    </div>
  );
};

export default Sidebar;
