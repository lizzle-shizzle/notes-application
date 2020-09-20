import React from 'react';
import { Col, Label, Row } from "reactstrap";

class Notes extends React.Component {
    constructor(props) {
        this.state = {
            noteText: ""
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        this.setState({ noteText: event.target.value });
    }

    handleSubmit = () => {
        let headers = new Headers();
        headers.append("Accept", "application/json");
        headers.append("Content-Type", "application/json");

        let data = {
            text: this.state.noteText
        }

        fetch(
            `${process.env.REACT_APP_NOTES_API_URL}/notes`,
            {
                method: "POST",
                headers: headers,
                body: JSON.stringify(data)
            }
        )
            .then(resp => {
                console.log(resp);
                if (resp.ok) {
                    resp
                        .json()
                        .then(json => {
                            alert("Successfully created note");
                            this.setState({
                                noteText: ""
                            })
                        })
                        .catch(err => {
                            console.log("err " + err);
                        });
                    return resp;
                } else {
                    resp.text().then(text => {
                        alert(`Request rejected with status ${resp.status} - ${text}`);
                    });
                }
            })
            .catch(err => {
                console.log(err);
            });
    }

    render() {
        return (
            <div>
                <Row>
                    <Col xs="12">
                        <label style={{ display: "block" }}>
                            Note text:
                    </label>
                        <textarea
                            value={this.state.noteText}
                            onChange={this.handleChange}
                            rows={10}
                            cols={35}
                            style={{ display: "block", marginLeft: "40%", marginBottom: "1%", padding: "1%" }}
                        />
                        <button
                            onClick={this.handleSubmit}
                            style={{ marginBottom: "2%" }}
                        >
                            Create note
                    </button>
                    </Col>
                </Row>
            </div>
        );
    }

}

export default Notes;
