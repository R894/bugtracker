import DashboardDisplay from '@/components/DashboardDisplay/DashboardDisplay'
import Layout from '@/layouts/layout'
import Head from 'next/head'

export default function Dashboard() {
  return (
    <>
      <Head>
        <title>Dashboard</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <Layout>
          <DashboardDisplay />
        </Layout>
      </main>
    </>
  )
}
