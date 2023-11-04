import { useRef } from "react";

export default function StateForm() {
  const name = useRef<HTMLInputElement>(null);
  const age = useRef<HTMLInputElement>(null);

  const show = () => {
    console.log(`Hello, ${name.current!.value} ${age.current!.value}`)
  }

  return (
    <>
      <form>
        <div>
          <label htmlFor="name">Name: </label>
          <input id="name" name="name" type="text"
            ref={name} defaultValue="user-name" />
        </div>
        <div>
          <label htmlFor="age">Age: </label>
          <input id="age" name="age" type="number"
            ref={age} defaultValue="18" />
        </div>
        <div>
          <button type="button" onClick={show}>
            send
          </button>
        </div>
      </form>
    </>
  )
}
