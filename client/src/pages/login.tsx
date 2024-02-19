import Footer from '@/components/Footer/Footer'
import LoginForm from '@/components/LoginForm/LoginForm'
import NavBar from '@/components/NavBar/NavBar'
import Head from 'next/head'

export default function Login() {
  return (
    <>
      <Head>
        <title>Login</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <NavBar />
        <div className="layout">
          <LoginForm />
        </div>
        <Footer />
      </main>
    </>
  )
}
