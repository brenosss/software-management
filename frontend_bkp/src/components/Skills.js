import { useState, useEffect } from 'react';
import { Row, Col, Card } from 'react-bootstrap';

import { getSkillsInfo } from '../services/skillsInfo'

export function SkillsList() {

  const [skillsInfo, setSkillsInfo] = useState([])
  const [skillTypes, setSkillTypes] = useState([])

  useEffect(() => {
    async function fetchData() {
      const responseSkillInfo = await getSkillsInfo()
      setSkillsInfo(responseSkillInfo.data)
      const uniqueSkillTypes = [...new Set(responseSkillInfo.data.map(item => item.Type))]
      setSkillTypes(uniqueSkillTypes)
    }
    fetchData()
  }, [])

  return (
    <Row>
      {
        skillTypes.map((type, index) => {
          return <SkillSection
              skills={skillsInfo.filter(skill => skill.Type === type)}
              type={type}
              key={index}
            />
        })
      }
    </Row>
  )
}

function SkillSection({skills, type}){
  return (
    <Row>
      <Col md={12}>
      <h1>{type}</h1>
      </Col>
      {
        skills.map((skill, index) => {
          return (
            <Col>
              <SkillItem skill={skill} key={index}></SkillItem>
            </Col>
          )
        })
      }
    </Row>
  )
}

function SkillItem({skill}) {
  return (
    <Card style={{ width: '18rem' }}>
      <Card.Body>
        <Card.Title>{skill.Name}</Card.Title>
        <Card.Subtitle className="mb-2 text-muted">
          Type: {skill.Type}
        </Card.Subtitle>
      </Card.Body>
    </Card>
  )
}