export type test = {
    message: string
}

export type post = {
    post_id: number;
    post_title: string;
    post_date: number;
    post_content: string;
    post_tags: tag[]
}

export type tag = {
    id:number;
    name:string;
}
