if __name__ == "__main__":
    import datetime
    import json
    import os
    import random

    import dotenv
    from mattermostdriver import Driver
    from mattermostdriver import exceptions as me

    # Load environment variables from .env as per 12-factor apps
    dotenv.load_dotenv()

    random.seed()

    MATCHED_MSG = "Hello! This week you have been matched up as conversation partners! I hope you meet up and have a great time :)"
    UNMATCHED_MSG = "Sorry! :( This week you haven't been matched with anyone. Better luck next week!"

    bot_username = os.getenv("ROBONONA_MATTERMOST_USERNAME")

    mattermost = Driver(
        {
            "url": os.getenv("ROBONONA_MATTERMOST_URL"),
            "login_id": bot_username,
            "password": os.getenv("ROBONONA_MATTERMOST_PASSWORD"),
            "port": 443,
        }
    )
    mattermost.login()

    try:
        mattermost_channel = mattermost.channels.get_channel_by_name_and_team_name(
            os.getenv("ROBONONA_MATTERMOST_TEAM"),
            os.getenv("ROBONONA_MATTERMOST_CHANNEL"),
        )

        channel_id = mattermost_channel['id']
        team_id = mattermost_channel['team_id']
    except me.ResourceNotFound:
        print("ERROR: mattermost team or channel name cannot be found")
        exit(1)

    page = 0
    users = []
    response = mattermost.users.get_users(
        params={
            'active': True,
            'in_channel': channel_id,
            'per_page': 200,
            'page': page,
        }
    )
    users.extend(response)
    while len(response) > 0:
        page += 1
        response = mattermost.users.get_users(
            params={
                'active': True,
                'in_channel': channel_id,
                'per_page': 200,
                'page': page,
            }
        )
        users.extend(response)

    users = [ (u['id'], u['username']) for u in users if u['username'] != bot_username ]

    random.shuffle(users)

    unmatched_user = None;
    if len(users) % 2 != 0:
        unmatched_user = users.pop()

    users = [users[i:i + 2] for i in range(0, len(users), 2)]

    if not os.getenv("ROBONONA_DRY_RUN") or os.getenv("ROBONONA_DRY_RUN") != "0":
        for pair in users:
            for pair in users:
                channel_id = mattermost.channels.create_group_channel(
                    pair[0][0],
                    pair[1][0]
                ])
                mattermost.posts.create_post(
                    channel_id,
                    MATCHED_MSG
                )

        if unmatched_user:
            channel_id = mattermost.channels.create_direct([user])
            mattermost.posts.create_post(
                channel_id,
                UNMATCHED_MSG
            )

    response = {
        'date': datetime.datetime.now().isoformat(),
        'dry_run': os.getenv("ROBONONA_DRY_RUN", False),
        'matched': [ (u[0][1], u[1][1]) for u in users ],
        'unmatched': unmatched_user,
    }

    print(json.dumps(response))
