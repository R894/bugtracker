import { useEffect, useState } from 'react'
import BugTable from './BugTable'
import ProjectsMenu from './ProjectsMenu'
import { useUserContext } from '@/context/UserContext'
import { getBugsByProjectId } from '@/api/bugService'
import { Bug } from '@/api/models/bug'
import CreateProjectModal from './CreateProjectModal'
import { Stack } from '@mui/material'
import CreateBugModal from './CreateBugModal'
import { useRouter } from 'next/router'
import { useProjectContext } from '@/context/ProjectContext'

const DashboardDisplay = ({projectId}:{projectId?: string}) => {
  const {
    user,
    token,
    logoutUser,
  } = useUserContext()
  const {selectedProject, setCurrentProject, updateUserProjects, getProjects} = useProjectContext()
  const [bugs, setBugs] = useState<Bug[]>([])
  const router = useRouter()

  useEffect(() => {
    updateUserProjects()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(()=> {
    if (!user) {
      router.push('/login')
    }
  },[user, router])

  useEffect(() => {
    if (!projectId) {
      const cachedProject = localStorage.getItem('project')
      if (cachedProject) {
        setCurrentProject(JSON.parse(cachedProject))
      }
    }
  }, [projectId, setCurrentProject])

  useEffect(() => {
    if (!token) return
    if (!selectedProject) return
    const getBugs = async () => {
      const response = await getBugsByProjectId(projectId? projectId: selectedProject.id, token)
      if (response?.code == 401) {
        logoutUser()
        return
      }
      setBugs(response)
    }
    getBugs()
  }, [logoutUser, projectId, selectedProject, token])

  return (
    <>
      <div style={{ padding: '16px' }}>
        <Stack direction={'row'}>
          <ProjectsMenu
            buttonText={
              selectedProject ? selectedProject.name : 'Select Project'
            }
            projects={user ? getProjects() : []}
            stateChanger={setCurrentProject}
          />
          <CreateProjectModal />
          {selectedProject ? (
            <CreateBugModal projectId={selectedProject.id} />
          ) : null}
        </Stack>
        <BugTable bugs={bugs} />
      </div>
    </>
  )
}

export default DashboardDisplay
