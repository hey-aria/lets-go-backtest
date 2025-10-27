import { useState } from 'preact/hooks';
import { h } from "preact";
import axios from "axios";
import styled from 'styled-components';

// Simple way to send csv files to our server
export default function DataUploader(props: any) {
    const [csv, setCsv] = useState<Blob>();
    const formData = new FormData();
    if (csv) {
        formData.append('file', csv);
    }

    function handleChange(e: any) {
        if (e.currentTarget.files) {
            setCsv(e.currentTarget.files[0]);
        }
    }

    function handleSubmit(e: any) {
        e.preventDefault();

        async function sendData() {
            const res = await axios.post('http://localhost:8080/upload', formData);
            console.log(res);
        }

        sendData();
    }

    return (
        <Wrapper>
            <Form onSubmit={handleSubmit}>
                <input type="file" accept=".csv" onChange={handleChange} />
                <button type="submit">Submit</button>
            </Form>
        </Wrapper>
    );
}

const Wrapper = styled.div`
    display: flex;
`;

const Form = styled.form`
    display: flex;
    flex-direction: column;
`;
