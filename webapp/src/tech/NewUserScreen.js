import React, { useState, useEffect } from "react";
import Modal from "react-modal";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import axios from "axios";
import "./NewUserScreen.css";

const customStyles = {
    content: {
        width: "60%",
        margin: "auto",
        padding: "20px"
    }
};

export function NewUserScreen({ show, onHide }) {

    // Initial user state
    const initialUser = {
        name: '',
        mail: '',
        age: 0
    };

    // State for validation messages
    const [validation, setValidation] = useState({
        name: initialUser && initialUser.name ? "" : "This field is required",
        mail: initialUser && initialUser.mail ? "" : "This field is required",
        age: initialUser && initialUser.age ? "" : "This field is required"
    });

    // State for updated user data
    const [updatedUser, setUpdatedUser] = useState({
        name: initialUser ? initialUser.name : "",
        mail: initialUser ? initialUser.mail : "",
        age: initialUser ? initialUser.age : ""
    });

    // Update validation messages when user data changes
    useEffect(() => {
        setValidation({
            name: updatedUser.name.trim() === "" ? "This field is required" : "",
            mail: updatedUser.mail.trim() === "" ? "This field is required" : "",
            age: updatedUser.age.toString().trim() === "" ? "This field is required" : ""
        });
    }, [updatedUser]);

    // Handle input changes
    const handleChange = (e) => {
        const { name, value } = e.target;

        // Handle different field types appropriately
        const updatedValue = name === "age" ? (value === "" ? "" : parseInt(value)) : value;

        // Update the state with the new value
        setUpdatedUser((prevUser) => ({
            ...prevUser,
            [name]: updatedValue,
        }));
    };

    // Handle onBlur event for input fields
    const handleBlur = (event) => {
        const { name, value } = event.target;
        setValidation((prevValidation) => ({
            ...prevValidation,
            [name]: value.trim() === "" ? "This field is required" : "",
        }));
    };

    // Function to add a new user
    const addUser = async () => {
        try {
            const response = await axios.post(`${process.env.REACT_APP_API_URL}/api/add-user`, updatedUser);
            console.log("User added:", response.data);
        } catch (error) {
            console.error("Error while adding user:", error);
        }
    };

    // Reload the main view
    const reloadMainView = () => {
        window.location.reload();
    };

    // Handle the "Add" button click
    const handleAddClick = () => {
        addUser();
        setTimeout(() => {
            onHide();
            reloadMainView();
        }, 500);
    };

    // Determine if the "Add" button should be disabled
    const isAddButtonDisabled = validation.name !== "" || validation.mail !== "" || validation.age !== "";

    return (
        <Modal
            isOpen={show}
            onRequestClose={onHide}
            contentLabel="Add User Modal"
            style={customStyles}
        >
            <div className="edit-form">
                <h2>Add User</h2>
                <Form>
                    <table className="table">
                        <tbody>
                        <tr>
                            <td><Form.Label>Name</Form.Label></td>
                            <td>
                                <Form.Control
                                    type="text"
                                    name="name"
                                    value={updatedUser.name}
                                    onChange={handleChange}
                                    onBlur={handleBlur}
                                    isInvalid={validation.name !== ""}
                                />
                                <Form.Control.Feedback type="invalid">
                                    {validation.name}
                                </Form.Control.Feedback>
                            </td>
                        </tr>
                        <tr>
                            <td><Form.Label>Email</Form.Label></td>
                            <td>
                                <Form.Control
                                    type="email"
                                    name="mail"
                                    value={updatedUser.mail}
                                    onChange={handleChange}
                                    onBlur={handleBlur}
                                    isInvalid={validation.mail !== ""}
                                />
                                <Form.Control.Feedback type="invalid">
                                    {validation.mail}
                                </Form.Control.Feedback>
                            </td>
                        </tr>
                        <tr>
                            <td><Form.Label>Age</Form.Label></td>
                            <td>
                                <Form.Control
                                    type="number"
                                    name="age"
                                    value={updatedUser.age}
                                    onChange={handleChange}
                                    onBlur={handleBlur}
                                    isInvalid={validation.age !== ""}
                                />
                                <Form.Control.Feedback type="invalid">
                                    {validation.age}
                                </Form.Control.Feedback>
                            </td>
                        </tr>
                        </tbody>
                    </table>
                </Form>
                <div className="edit-buttons">
                    <Button variant="secondary" onClick={onHide}>
                        Close
                    </Button>
                    <Button
                        variant="primary"
                        onClick={handleAddClick}
                        disabled={isAddButtonDisabled}
                    >
                        Add
                    </Button>
                </div>
            </div>
        </Modal>
    );
}
