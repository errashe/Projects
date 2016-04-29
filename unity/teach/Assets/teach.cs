using UnityEngine;
using System;
using System.Collections;
using System.Collections.Generic;

public class teach : MonoBehaviour {

	public float start;
	public float end;
	public float step;

	public float width;
	public float height;

	public GameObject spawn;
	List<GameObject> cubes;
	List<Vector3> vectors;

	private float wx = 1f;
	private float dx = 1f;
	private float wy = 1f;
	private float dy = 1f;
	private float a1 = 1f;
	private float b1 = 1f;

	private float delta = .05f;

	void generateVectors() {
		vectors.Clear ();

		float a = a1 / Mathf.Min(a1, b1);
		float b = b1 / Mathf.Min (a1, b1);

		float itlen;
		if (wx + wy > 30) {
			itlen = 0.005f;
		} else if (wx + wy > 12) {
			itlen = 0.01f;
		} else {
			itlen = 0.02f;
		}

		float xlast = (Mathf.Cos (wx * 0 + dx) * (width / 2 - 10) / b + width / 2);
		float ylast = height - (Mathf.Cos (wy * 0 + dy) * (height / 2 - 10) / a + height / 2);

		for(float i = 0; i < 6.4f; i += itlen){
			float x = (Mathf.Cos (wx * i + dx) * (width / 2 - 10) / b + width / 2);
			float y = height - (Mathf.Cos (wy * i + dy) * (height / 2 - 10) / a + height / 2);
			vectors.Add (new Vector3 (x, y));
			xlast = x;
			ylast = y;
		}
	}

	void generateLines() {
		LineRenderer lr = GetComponent<LineRenderer> ();
		lr.SetColors (Color.black, Color.black);
		lr.SetVertexCount (vectors.Count);
		lr.SetWidth (.1f, .1f);
		lr.SetPositions (vectors.ToArray ());
	}

	// Use this for initialization
	void Start () {
		
		vectors = new List<Vector3> ();
		generateVectors ();
		generateLines ();
	}

	// Update is called once per frame
	void Update () {

		if (Input.GetKey (KeyCode.Escape)) {
			#if UNITY_EDITOR
			UnityEditor.EditorApplication.isPlaying = false;
			#else
			Application.Quit();
			#endif
		}

		if (Input.GetKey (KeyCode.T)) { wx += delta; }
		if (Input.GetKey (KeyCode.G)) { wx -= delta; }

		if (Input.GetKey (KeyCode.Y)) { dx += delta; }
		if (Input.GetKey (KeyCode.H)) { dx -= delta; }

		if (Input.GetKey (KeyCode.U)) { wy += delta; }
		if (Input.GetKey (KeyCode.J)) { wy -= delta; }

		if (Input.GetKey (KeyCode.I)) { dy += delta; }
		if (Input.GetKey (KeyCode.K)) { dy -= delta; }

		if (Input.GetKey (KeyCode.O)) { a1 += delta; }
		if (Input.GetKey (KeyCode.L)) { a1 -= delta; }

		if (Input.GetKey (KeyCode.P)) { b1 += delta; }
		if (Input.GetKey (KeyCode.Semicolon)) { b1 -= delta; }

		someClicked ();

	}

	List<KeyCode> codes = new List<KeyCode>(new KeyCode[]{
		KeyCode.T, KeyCode.G, 
		KeyCode.Y, KeyCode.H, 
		KeyCode.U, KeyCode.J, 
		KeyCode.I, KeyCode.K, 
		KeyCode.O, KeyCode.L,
		KeyCode.P, KeyCode.Semicolon});

	void someClicked() {
		foreach (KeyCode code in codes) {
			if (Input.GetKey (code)) {
				generateVectors ();
				generateLines ();
			}
		}
	}

	void OnGUI() {
		GUI.Label (new Rect (25, 25, 100, 100), "T/G | WX: " + Math.Round(wx, 2).ToString ());
		GUI.Label (new Rect (25, 45, 100, 100), "Y/H | DX: " + Math.Round (dx, 2).ToString ());
		GUI.Label (new Rect (25, 65, 100, 100), "U/J | WY: " + Math.Round (wy, 2).ToString ());
		GUI.Label (new Rect (25, 85, 100, 100), "I/K | DY: " + Math.Round (dy, 2).ToString ());
		GUI.Label (new Rect (25, 105, 100, 100), "O/L | A1: " + Math.Round (a1, 2).ToString ());
		GUI.Label (new Rect (25, 125, 100, 100), "P/; | B1: " + Math.Round (b1, 2).ToString ());
	}
}
