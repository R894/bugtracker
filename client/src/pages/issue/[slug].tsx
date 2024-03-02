import BugForm from '@/components/BugForm/BugForm'
import CommentForm from '@/components/Comments/CommentForm'
import Layout from '@/layouts/Layout'
import Head from 'next/head'
import { useRouter } from 'next/router'

export default function Home() {
  const router = useRouter()
  return (
    <>
      <Head>
        <title>Bugtracker</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <Layout>
          {router.query.slug && <BugForm bugId={router.query.slug as string} />}
          <CommentForm/>
        </Layout>
      </main>
    </>
  )
}
