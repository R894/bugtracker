import RegisterForm from '@/components/RegisterForm/RegisterForm'
import Layout from '@/layouts/Layout'
import Head from 'next/head'

export default function Login() {
  return (
    <>
      <Head>
        <title>Login</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <Layout>
          <RegisterForm />
        </Layout>
      </main>
    </>
  )
}
