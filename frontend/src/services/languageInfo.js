import axios from 'axios'

export async function getLanguageInfo() {
    const response = await axios.get('http://localhost:8080/languages/info')
    return response.data
}