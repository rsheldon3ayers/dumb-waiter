import Link from 'next/link'

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      {/* todo: add login form */}
      <div className="flex flex-col items-center">
        <h1 className="text-4xl font-bold">Dumb Waiter</h1>
        <Link className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800" href="/dashboard">Dashboard</Link>
        </div>
    </main>
  )
}
