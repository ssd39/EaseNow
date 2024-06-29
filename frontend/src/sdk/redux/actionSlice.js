import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import {
  ApolloClient,
  InMemoryCache,
  gql,
} from "@apollo/client";

const client = new ApolloClient({
  uri: "https://api.studio.thegraph.com/query/76928/easenow/v0.1",
  cache: new InMemoryCache(),
});

const fetchUser = createAsyncThunk(
  "users/fetchUser",
  async (userAddress, thunkAPI) => {
    const result = await client.query({
      query: gql`
      query MyData {
        userRegistreds(where: {userAddress: "${userAddress}"}){
          privateData,
          creditLimit,
        }
      }
      `,
    });
    if(result.data.userRegistreds.length <= 0) {
      return {
        isNewUser: true,
        creditLimit: 0
      }
    }
    return { isNewUser: false, creditLimit: result.data.userRegistreds[0].creditLimit};
  }
);

export const actionSlice = createSlice({
  name: "action",
  initialState: {
    isNewUser: true,
    isUserFetched: false,
    creditLimit: 0,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchUser.fulfilled, (state, action) => {
      state.isNewUser = action.payload.isNewUser;
      state.creditLimit = action.payload.creditLimit;
      state.isUserFetched = true;
    });
    builder.addCase(fetchUser.rejected, (state, action) => {
      console.error("Error while fetching new user!", action.error);
    });
  },
});

// Action creators are generated for each case reducer function
export const {} = actionSlice.actions;
export { fetchUser }
export default actionSlice.reducer;
