using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;
using System.Net;
using System.Threading;

namespace package
{
    public partial class FormMain : Form
    {
        public class T_PackInfo
        {
            public int Number { get; set; }
            public string Name { get; set; }
            public string Path { get; set; }
        }

        public class T_UnpackInfo
        {
            public int Number { get; set; }
            public string Name { get; set; }
            public string Size { get; set; }
            public string Type { get; set; }
        }

        private Thread m_tPack;
        private Thread m_tUnpack;
        private List<T_PackInfo> m_vecPackInfo = new List<T_PackInfo>();
        private List<T_UnpackInfo> m_vecUnpackInfo = new List<T_UnpackInfo>();

        public FormMain()
        {
            InitializeComponent();
        }

        private void FormMain_Load(object sender, EventArgs e)
        {
            // FormMain Initial
            AutoScaleMode = AutoScaleMode.Dpi;
            MaximizeBox = false;
            MinimizeBox = true;
            DoubleBuffered = true;
            StartPosition = FormStartPosition.CenterScreen;

            // TabPack Initial
            listView_pack.View = View.Details;
            listView_pack.CheckBoxes = true;
            listView_pack.Columns.Add("number", 60, HorizontalAlignment.Center);
            listView_pack.Columns.Add("files", 240, HorizontalAlignment.Center);
            listView_pack.Columns.Add("detials", 400, HorizontalAlignment.Center);

            comboBox_pack.DropDownStyle = ComboBoxStyle.DropDownList;
            comboBox_pack.Items.Clear();
            comboBox_pack.Items.Add("aes");
            comboBox_pack.Items.Add("3des");
            comboBox_pack.Items.Add("des");
            comboBox_pack.Items.Add("rsa");
            comboBox_pack.Items.Add("base64");
            comboBox_pack.SelectedIndex = 0;

            textBox_pack_path.ReadOnly = true;

            timer_pack_process.Enabled = true;
            timer_pack_process.Stop();

            // TabUnpack Initial
            listView_unpack.View = View.Details;
            listView_unpack.Columns.Add("number", 60, HorizontalAlignment.Center);
            listView_unpack.Columns.Add("files", 240, HorizontalAlignment.Center);
            listView_unpack.Columns.Add("size", 200, HorizontalAlignment.Center);
            listView_unpack.Columns.Add("type", 200, HorizontalAlignment.Center);

            textBox_unpack_src.ReadOnly = true;
            textBox_unpack_dest.ReadOnly = true;
        }

        private void FormMain_FormClosed(object sender, FormClosedEventArgs e)
        {
            
        }

        private void Update_pack_listView()
        {
            listView_pack.Items.Clear();
            listView_pack.BeginUpdate();
            for (int i = 0; i < m_vecPackInfo.Count; ++i)
            {
                ListViewItem item = new ListViewItem();
                item.Text = m_vecPackInfo[i].Number.ToString();
                item.SubItems.Add(m_vecPackInfo[i].Name);
                item.SubItems.Add(m_vecPackInfo[i].Path);
                listView_pack.Items.Add(item);
            }
            listView_pack.EndUpdate();
        }

        private string Get_pack_json()
        {
            // check src file list
            if (m_vecPackInfo.Count <= 0)
            {
                MessageBox.Show("Please select src files name!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }
            // check dest file
            if (textBox_pack_path.Text == "")
            {
                MessageBox.Show("Please select dest file path!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }

            string src = "";
            for (int i = 0; i < m_vecPackInfo.Count - 1; ++i)
            {
                src += string.Format("\"{0}\",", m_vecPackInfo[i].Path);
            }
            src += string.Format("\"{0}\"", m_vecPackInfo[m_vecPackInfo.Count - 1].Path);

            string dest = string.Format("\"{0}\"", textBox_pack_path.Text);

            string type = "";
            switch (comboBox_pack.SelectedIndex)
            {
                case 0:
                    type = "\"aes\"";
                    break;
                case 1:
                    type = "\"3des\"";
                    break;
                case 2:
                    type = "\"des\"";
                    break;
                case 3:
                    type = "\"rsa\"";
                    break;
                case 4:
                    type = "\"base64\"";
                    break;
                default:
                    type = "\"aes\"";
                    break;
            }

            string body = "";
            body += "{";
            body += string.Format("\"src\":[{0}],\"dest\":{1},\"type\":{2}", src, dest, type);
            body += "}";
            body = body.Replace('\\', '/');
            return body;
        }

        private string Get_unpack_json()
        {
            // check src file
            if (textBox_unpack_src.Text == "")
            {
                MessageBox.Show("Please select src file name!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }
            // check dest path
            if (textBox_unpack_dest.Text == "")
            {
                MessageBox.Show("Please select dest file path!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }

            string src = string.Format("\"{0}\"", textBox_unpack_src.Text);
            string dest = string.Format("\"{0}\"", textBox_unpack_dest.Text);

            string body = "";
            body += "{";
            body += string.Format("\"src\":{0},\"dest\":{1}", src, dest);
            body += "}";
            body = body.Replace('\\', '/');
            return body;
        }

        private void PackThread(object sender)
        {
            string body = sender as string;
            try
            {
                byte[] requestBody = Encoding.ASCII.GetBytes(body);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/pack");
                request.Method = "POST";
                request.ProtocolVersion = HttpVersion.Version11;
                request.ContentType = "application/json";
                request.ContentLength = requestBody.Length;
                using (Stream reqStream = request.GetRequestStream())
                {
                    reqStream.Write(requestBody, 0, requestBody.Length);
                }
                using (HttpWebResponse response = (HttpWebResponse)request.GetResponse())
                {
                    if (response.StatusCode == HttpStatusCode.OK)
                    {
                        MessageBox.Show("pack success!", "Information", MessageBoxButtons.OK, MessageBoxIcon.Information);
                    }
                    else
                    {
                        MessageBox.Show("Response status code not right!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                    }
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
            }
        }

        private void UnpackThread(object sender)
        {
            string body = sender as string;
            try
            {
                byte[] requestBody = Encoding.ASCII.GetBytes(body);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/unpack");
                request.Method = "POST";
                request.ProtocolVersion = HttpVersion.Version11;
                request.ContentType = "application/json";
                request.ContentLength = requestBody.Length;
                using (Stream reqStream = request.GetRequestStream())
                {
                    reqStream.Write(requestBody, 0, requestBody.Length);
                }
                using (HttpWebResponse response = (HttpWebResponse)request.GetResponse())
                {
                    if (response.StatusCode == HttpStatusCode.OK)
                    {
                        MessageBox.Show("unpack success!", "Information", MessageBoxButtons.OK, MessageBoxIcon.Information);
                    }
                    else
                    {
                        MessageBox.Show("Response status code not right!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                    }
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
            }
        }

        private void Button_pack_add_Click(object sender, EventArgs e)
        {
            OpenFileDialog dialog = new OpenFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.Multiselect = true;
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                foreach (var i in dialog.FileNames)
                {
                    bool flag = false;
                    foreach (var j in m_vecPackInfo)
                    {
                        if(j.Path == i)
                        {
                            flag = true;
                            break;
                        }
                    }
                    if(!flag)
                    {
                        T_PackInfo packInfo = new T_PackInfo();
                        FileInfo file = new FileInfo(i);
                        packInfo.Number = m_vecPackInfo.Count + 1;
                        packInfo.Name = file.Name;
                        packInfo.Path = i;
                        m_vecPackInfo.Add(packInfo);
                    }
                }

                Update_pack_listView();
            }

        }

        private void Button_pack_delete_Click(object sender, EventArgs e)
        {
            if(listView_pack.CheckedItems.Count > 0)
            {
                for (int i = 0; i < listView_pack.CheckedItems.Count; ++i)
                {
                    for (int j = 0; j < m_vecPackInfo.Count;)
                    {
                        if(m_vecPackInfo[j].Number.ToString() == listView_pack.CheckedItems[i].Text)
                        {
                            m_vecPackInfo.RemoveAt(j);
                        }
                        else
                        {
                            ++j;
                        }
                    }
                }

                for(int i = 0; i < m_vecPackInfo.Count; ++i)
                {
                    m_vecPackInfo[i].Number = i + 1;
                }

                Update_pack_listView();
            }
        }

        private void Button_pack_select_Click(object sender, EventArgs e)
        {
            SaveFileDialog dialog = new SaveFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.FileName = "undefine.pak";
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                textBox_pack_path.Text = dialog.FileName;
            }

        }

        private void Button_pack_execute_Click(object sender, EventArgs e)
        {
            string body = Get_pack_json();
            if (body == "")
                return;

            m_tPack = new Thread(new ParameterizedThreadStart(PackThread));
            m_tPack.IsBackground = true;
            m_tPack.Start(body);
        }

        private void Timer_pack_process_Tick(object sender, EventArgs e)
        {
            try
            {
                string src = "";
                string type = "";

                for (int i = 0; i < m_vecPackInfo.Count - 1; ++i)
                {
                    src += string.Format("\"{0}\",", m_vecPackInfo[i].Path);
                }
                src += string.Format("\"{0}\"", m_vecPackInfo[m_vecPackInfo.Count - 1].Path);

                switch (comboBox_pack.SelectedIndex)
                {
                    case 0:
                        type = "\"aes\"";
                        break;
                    case 1:
                        type = "\"3des\"";
                        break;
                    case 2:
                        type = "\"des\"";
                        break;
                    case 3:
                        type = "\"rsa\"";
                        break;
                    case 4:
                        type = "\"base64\"";
                        break;
                    default:
                        type = "\"aes\"";
                        break;
                }

                string body = "";
                body += "{";
                body += string.Format("\"src\":[{0}],\"type\":{1}", src, type);
                body += "}";
                body = body.Replace('\\', '/');

                byte[] requestBody = Encoding.ASCII.GetBytes(body);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/pack/p");
                request.Method = "POST";
                request.ProtocolVersion = HttpVersion.Version10;
                request.ContentType = "application/json";
                request.ContentLength = requestBody.Length;
                using (Stream reqStream = request.GetRequestStream())
                {
                    reqStream.Write(requestBody, 0, requestBody.Length);
                }
                using (HttpWebResponse response = (HttpWebResponse)request.GetResponse())
                {
                    //request.Abort();
                    if (response.StatusCode == HttpStatusCode.OK)
                    {
                        string responseContent = "";
                        using (Stream resStream = response.GetResponseStream())
                        {
                            using (StreamReader reader = new StreamReader(resStream, Encoding.UTF8))
                            {
                                responseContent = reader.ReadToEnd().ToString();
                            }
                        }
                    }
                    else
                    {
                        MessageBox.Show("Response status code not right!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                        timer_pack_process.Stop();
                    }
                    //response.Close();
                }
            }
            catch(Exception ex)
            {
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                timer_pack_process.Stop();
            }
        }

        private void Button_unpack_src_Click(object sender, EventArgs e)
        {
            OpenFileDialog dialog = new OpenFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.Multiselect = false;
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                textBox_unpack_src.Text = dialog.FileName;
            }
        }

        private void Button_unpack_verbose_Click(object sender, EventArgs e)
        {

        }

        private void Button_unpack_dest_Click(object sender, EventArgs e)
        {
            SaveFileDialog dialog = new SaveFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.FileName = "undefine";
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                textBox_unpack_dest.Text = Path.GetDirectoryName(dialog.FileName) + "\\";
            }
        }

        private void Button_unpack_execute_Click(object sender, EventArgs e)
        {
            string body = Get_unpack_json();
            if (body == "")
                return;

            m_tUnpack = new Thread(new ParameterizedThreadStart(UnpackThread));
            m_tUnpack.IsBackground = true;
            m_tUnpack.Start(body);
        }
    }
}
