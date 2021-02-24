import { useState, useEffect } from 'react';
import { Row, Col, Card } from 'react-bootstrap';

import { getLanguageInfo } from '../services/languageInfo'

export function LanguageList() {
    const [ languageInfo, setLanguagesInfo ] = useState([])

    useEffect(() => {
      async function fetchData(){
        const languagesInfo = await getLanguageInfo()
        setLanguagesInfo(languagesInfo.data)
      }
      fetchData()
    }, [])

    return (
      <Row>
        {
          languageInfo.map((language) => {
            return (
              <Col>
                <LanguageItem language={language}></LanguageItem>
              </Col>
            )
          })
        }
      </Row>
    )
  }
  
  function LanguageItem({language}) {
    return (
      <Card style={{ width: '18rem' }}>
        <Card.Body>
          <Card.Title>{language.name}</Card.Title>
          <Card.Subtitle className="mb-2 text-muted">
            Popularity: {language.popularity}
          </Card.Subtitle>
          <Card.Subtitle className="mb-2 text-muted">
            Learning: {language.learning}
          </Card.Subtitle>
          {false && language.use.map((use) => <Card.Link href="#">{use}</Card.Link>)}
        </Card.Body>
      </Card>
    )
  }