import axios from 'axios'

export async function getSkillsInfo() {
    const response = await axios.get('http://localhost:8080/skills')
    return response.data
}