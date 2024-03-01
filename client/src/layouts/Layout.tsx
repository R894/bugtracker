import Footer from '@/components/Footer/Footer'
import NavBar from '@/components/NavBar/NavBar'
import { type ReactNode } from 'react'

const Layout = ({ children }: { children: ReactNode }) => {
  return (
    <>
      <NavBar />
      <div className="layout">{children}</div>
      <Footer />
    </>
  )
}

export default Layout
