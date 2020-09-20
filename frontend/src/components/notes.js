import React from 'react';
import { Col, Label, Row } from "reactstrap";

class Notes extends React.Component {
    constructor(props) {
        this.state = {
            notes: [],
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
                            this.fetchNotes()
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

    fetchNotes() {
        fetch(
            `${process.env.REACT_APP_NOTES_API_URL}/notes`
        )
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else {
                    console.log("error");
                    console.log(response);
                }
            })
            .then(json => {
                this.setState({
                    notes: json
                });
            });
    }

    componentWillMount() {
        this.fetchNotes()
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
                <Row>
                    <Col size={{ xs: 12 }}>
                        <label>
                            <b>My notes:</b>
                        </label>
                        <table style={{ width: "50%", marginLeft: "35%", textAlign: "left", borderSpacing: "1em 1em" }}>
                            {this.state.notes !== undefined && this.state.notes.map((note, index) => {
                                const { id, text, created_timestamp } = note
                                return (
                                    <tr key={id}>
                                        <td></td>
                                        <td style={{ width: "15%" }}>{created_timestamp.substring(0, 10)}</td>
                                        <td>{text}</td>
                                    </tr>
                                )
                            })}
                        </table>
                    </Col>
                </Row>
            </div>
        );
    }

}

export default Notes;
