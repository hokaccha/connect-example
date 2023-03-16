import type { GetServerSideProps } from "next";
import { FC, useEffect, useState } from "react";
import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { GreetService } from "../gen/greet/v1/greet_connect";

const transport = createConnectTransport({ baseUrl: "http://localhost:8080" });
const client = createPromiseClient(GreetService, transport);

async function greet(name: string): Promise<string> {
  const response = await client.greet({ name });
  return response.greeting;
}

type Props = {
  name: string;
  greeting: string;
};

const TopPage: FC<Props> = (props) => {
  const [name, setName] = useState(props.name);
  const [greeting, setGreeting] = useState(props.greeting);

  useEffect(() => {
    greet(name).then((g) => setGreeting(g));
  }, [name]);

  return (
    <div>
      <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
      <div>{greeting}</div>
    </div>
  );
};

export const getServerSideProps: GetServerSideProps<Props> = async (context) => {
  const name = context.query.name?.toString() || "John";
  const greeting = await greet(name);
  return { props: { name, greeting } };
};

export default TopPage;
