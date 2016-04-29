using UnityEngine;
using System.Collections;

public class camera : MonoBehaviour {

	private Vector3 mouseOrigin;
	private bool isRotating;

	void Start () {
		Cursor.lockState = CursorLockMode.Locked;
	}

	public float speedH = 2.0f;
	public float speedV = 2.0f;

	private float yaw = 0.0f;
	private float pitch = 0.0f;

	void Update () {
		if (Input.GetKey (KeyCode.A)) {
			transform.Translate (10 * Vector3.left * Time.deltaTime);
		}
		if (Input.GetKey (KeyCode.D)) {
			transform.Translate (10 * Vector3.right * Time.deltaTime);
		}
		if (Input.GetKey (KeyCode.W)) {
			transform.Translate (10 * Vector3.forward * Time.deltaTime);
		}
		if (Input.GetKey (KeyCode.S)) {
			transform.Translate (10 * Vector3.back * Time.deltaTime);
		}
		if (Input.GetKey (KeyCode.Space)) {
			transform.Translate (10 * Vector3.up * Time.deltaTime);
		}
		if (Input.GetKey (KeyCode.C)) {
			transform.Translate (10 * Vector3.down * Time.deltaTime);
		}

		yaw += speedH * Input.GetAxis("Mouse X");
		pitch -= speedV * Input.GetAxis("Mouse Y");

		transform.eulerAngles = new Vector3(pitch, yaw, 0.0f);

//		if (Input.GetKey (KeyCode.LeftControl)) {
//			if (Cursor.lockState == CursorLockMode.Locked) {
//				Cursor.lockState = CursorLockMode.None;
//				Cursor.visible = true;
//			} else {
//				Cursor.lockState = CursorLockMode.Locked;
//				Cursor.visible = false;
//			}
//		}
	}
}
