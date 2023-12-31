import Link from 'next/link'

export default function Dashboard() {
    return (
        <div className="flex flex-col items-center justify-center min-h-screen py-2">
            <Link href="/create-ticket">Add a ticket</Link>
            <Link href="/tickets">View tickets</Link>
            <Link href="/"></Link>
        </div>
    )
}