import ProjectProvider from '@/context/ProjectContext'
import UserProvider from '@/context/UserContext'
import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { Roboto } from 'next/font/google'

const roboto = Roboto({
  weight: ['300', '400', '500', '700'],
  style: ['normal', 'italic'],
  subsets: ['latin'],
  display: 'swap',
})

export default function App({ Component, pageProps }: AppProps) {
  return (
    <div className={roboto.className}>
      <UserProvider>
        <ProjectProvider>
          <Component {...pageProps} />
        </ProjectProvider>
      </UserProvider>
    </div>
  )
}
