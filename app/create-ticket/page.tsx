import Link from 'next/link';

export default function CreateTicket() {
    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <form className="flex flex-col items-center">
                <label className="text-2xl font-bold">Create a ticket</label>
                <input className="border-2 border-gray-400 rounded-lg p-2 m-2" type="text" placeholder="Title" />
                <input className="border-2 border-gray-400 rounded-lg p-2 m-2" type="text" placeholder="Description" />
                <input className="border-2 border-gray-400 rounded-lg p-2 m-2" type="text" placeholder="Assignee" />
                <button className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800" type="submit">Create</button>
            </form>
        </main>
    );
}