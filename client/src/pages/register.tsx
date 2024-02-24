import Footer from '@/components/Footer/Footer'
import NavBar from '@/components/NavBar/NavBar'
import RegisterForm from '@/components/RegisterForm/RegisterForm'
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
          <RegisterForm />
        </div>
        <Footer />
      </main>
    </>
  )
}