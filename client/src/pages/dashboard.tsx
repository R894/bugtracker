import DashboardDisplay from '@/components/DashboardDisplay/DashboardDisplay'
import Footer from '@/components/Footer/Footer'
import NavBar from '@/components/NavBar/NavBar'
import Head from 'next/head'

export default function Dashboard() {
  return (
    <>
      <Head>
        <title>Dashboard</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <NavBar />
        <div className="layout">
          <DashboardDisplay/>
        </div>
        <Footer />
      </main>
    </>
  )
}
