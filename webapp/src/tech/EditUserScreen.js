import React, { useEffect, useState } from "react";
import Modal from "react-modal";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import axios from "axios";
import "./EditUserScreen.css";

const customStyles = {
    content: {
        width: "60%",
        margin: "auto",
        padding: "20px",
    },
};

export function EditUserScreen({ show, onHide, user }) {
    // State for input validation messages
    const [validation, setValidation] = useState({
        name: user && user.name ? "" : "This field is required",
        mail: user && user.mail ? "" : "This field is required",
        age: user && user.age ? "" : "This field is required",
    });

    // State to store updated user data
    const [updatedUser, setUpdatedUser] = useState({
        id: user ? user.id:"",
        name: user ? user.name : "",
        mail: user ? user.mail : "",
        age: user ? user.age : "",
    });

    useEffect(() => {
        // Update validation messages based on input values
        setValidation({
            name: updatedUser.name === "" ? "This field is required" : "",
            mail: updatedUser.mail === "" ? "This field is required" : "",
            age: updatedUser.age === "" ? "This field is required" : "",
        });
    }, [updatedUser]);

    // Handle input change
    const handleChange = (e) => {
        const { name, value } = e.target;

        // Handle different field types appropriately
        const updatedValue =
            name === "age" ? (value === "" ? "" : parseInt(value)) : value;

        // Update the state with the new value
        setUpdatedUser((prevUser) => ({
            ...prevUser,
            [name]: updatedValue,
        }));
    };

    // Handle input blur event for validation
    const handleBlur = (event) => {
        const { name, value } = event.target;
        setValidation((prevValidation) => ({
            ...prevValidation,
            [name]: value.trim() === "" ? "This field is required" : "",
        }));
    };

    // Save user data
    const saveUser = async () => {
        try {
            const response = await axios.post(
                `${process.env.REACT_APP_API_URL}/api/update-user`,
                updatedUser
            );
            console.log("User updated:", response.data);
        } catch (error) {
            console.error("Error updating user:", error);
        }
    };

    // Reload the main view after saving
    const reloadMainView = () => {
        window.location.reload();
    };

    // Handle save button click
    const handleSaveClick = () => {
        saveUser();
        setTimeout(() => {
            onHide();
            reloadMainView();
        }, 500);
    };

    // Check if the add button should be disabled
    const isAddButtonDisabled =
        validation.name !== "" || validation.mail !== "" || validation.age !== "";

    return (
        <Modal
            isOpen={show}
            onRequestClose={onHide}
            contentLabel="Edit User Modal"
            style={customStyles}
        >
            <div className="edit-form">
                <h2>Edit User</h2>
                <Form>
                    <table className="table">
                        <tbody>
                        <tr>
                            <td>
                                <Form.Label>Name</Form.Label>
                            </td>
                            <td>
                                <Form.Control
                                    type="text"
                                    name="name"
                                    value={updatedUser.name}
                                    onChange={handleChange}
                                    onBlur={handleBlur}
                                    required
                                    isInvalid={validation.name !== ""}
                                />
                                <Form.Control.Feedback type="invalid">
                                    {validation.name}
                                </Form.Control.Feedback>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <Form.Label>Email</Form.Label>
                            </td>
                            <td>
                                <Form.Control
                                    type="email"
                                    name="mail"
                                    value={updatedUser.mail}
                                    onChange={handleChange}
                                    onBlur={handleBlur}
                                    required
                                    isInvalid={validation.mail !== ""}
                                />
                                <Form.Control.Feedback type="invalid">
                                    {validation.mail}
                                </Form.Control.Feedback>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <Form.Label>Age</Form.Label>
                            </td>
                            <td>
                                <Form.Control
                                    type="number"
                                    name="age"
                                    value={updatedUser.age}
                                    onChange={handleChange}
                                    onBlur={handleBlur}
                                    required
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
                        onClick={handleSaveClick}
                        disabled={isAddButtonDisabled}
                    >
                        Save
                    </Button>
                </div>
            </div>
        </Modal>
    );
}
