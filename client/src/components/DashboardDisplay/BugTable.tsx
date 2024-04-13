import { Bug } from '@/api/models/bug'
import { DataGrid, GridColDef } from '@mui/x-data-grid'
import moment from 'moment'
import Link from 'next/link'

const columns: GridColDef[] = [
  {
    field: 'title',
    headerName: 'Title',
    width: 300,
    renderCell: (params) => (
      <Link href={`/issue/${params.row.id}`}>{params.row.title}</Link>
    ),
  },
  { field: 'priority', headerName: 'priority' },
  { field: 'status', headerName: 'Status' },
  { field: 'asignee', headerName: 'Asignee' },
  {
    field: 'createdAt',
    headerName: 'Date',
    renderCell: (params) => {
      return moment(params.row.createdAt).format('DD/MM/YYYY')
    },
  },
]

const BugTable = ({ bugs }: { bugs: Bug[] }) => {
  return (
    <DataGrid rows={bugs || []} columns={columns} autoHeight />
  )
}

export default BugTable
