import React, { useState } from "react";

export default function FormTextarea() {
  const [form, setForm] = useState({
    comment: ""
  });

  const handleForm = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value
    })
  }

  const show = () => {
    console.log(`comment: ${form.comment}`)
  }

  return (
    <>
      <form>
        <div>
          <label htmlFor="comment">Comment: </label>
          <br />
          <textarea id="comment" name="comment"
          cols={30} rows={7}
          value={form.comment}
          onChange={handleForm}
          />
          <br />
          <button type="button" onClick={show}>
            send
          </button>
        </div>
      </form>
    </>
  )
}
