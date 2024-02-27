import { useContext, useEffect, useState } from 'react'
import BugList from './BugList'
import ProjectsMenu from './ProjectsMenu'
import { UserContext, UserContextType } from '@/context/UserContext'
import { getBugsByProjectId } from '@/api/bugService'
import { Bug } from '@/api/models/bug'

const DashboardDisplay = () => {
  const { updateUserProjects, user, token } = useContext(
    UserContext,
  ) as UserContextType
  const [selectedProjectId, setSelectedProjectId] = useState('')
  const [bugs, setBugs] = useState<Bug[]>([])

  useEffect(() => {
    updateUserProjects()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => {
    if (!token) return
    const getBugs = async () => {
      const response = await getBugsByProjectId(selectedProjectId, token)
      console.log("project id", selectedProjectId)
      console.log(response)
      setBugs(response)
    }
    getBugs()
  }, [selectedProjectId, token])

  return (
    <>
      <ProjectsMenu
        projects={user ? user.projects : []}
        stateChanger={setSelectedProjectId}
      />
      <BugList bugs={bugs} />
    </>
  )
}

export default DashboardDisplay
