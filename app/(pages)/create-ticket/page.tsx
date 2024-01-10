import Link from 'next/link';
import Description from '@/app/components/description';

export default function CreateTicket() {
    return (
        <main className="flex w-4/5 min-h-screen flex-col p-24">
            <form className="flex flex-col items-center basis-9/12">
                <label className="text-2xl font-bold">Create a ticket</label>
                <input className="border-2 border-gray-400 rounded-lg p-2 m-2" type="text" placeholder="Title" />
                <Description />
                <input className="border-2 border-gray-400 rounded-lg p-2 m-2" type="text" placeholder="Assignee" />
                <button className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800" type="submit">Create</button>
            </form>
        </main>
    );
}